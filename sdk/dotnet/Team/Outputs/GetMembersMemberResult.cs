// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Heroku.Team.Outputs
{

    [OutputType]
    public sealed class GetMembersMemberResult
    {
        /// <summary>
        /// Email address of the team member.
        /// </summary>
        public readonly string Email;
        /// <summary>
        /// Whether the user is federated and belongs to an Identity Provider.
        /// </summary>
        public readonly bool Federated;
        /// <summary>
        /// Role in the team.
        /// </summary>
        public readonly string Role;
        /// <summary>
        /// Unique identifier of the team member on the team.
        /// </summary>
        public readonly string TeamMemberId;
        /// <summary>
        /// Whether the Enterprise team member has two-factor authentication enabled.
        /// </summary>
        public readonly bool TwoFactorAuthentication;
        /// <summary>
        /// Unique identifier of the team member. This is the member's user ID in Heroku.
        /// </summary>
        public readonly string UserId;

        [OutputConstructor]
        private GetMembersMemberResult(
            string email,

            bool federated,

            string role,

            string teamMemberId,

            bool twoFactorAuthentication,

            string userId)
        {
            Email = email;
            Federated = federated;
            Role = role;
            TeamMemberId = teamMemberId;
            TwoFactorAuthentication = twoFactorAuthentication;
            UserId = userId;
        }
    }
}
